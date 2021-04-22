var Generator = require('yeoman-generator');
var fs = require('fs')
var sanitize = require("../sanitize")
const parseMakefile = require('@kba/makefile-parser')
let ora = require("ora");

module.exports = class extends Generator {

  // The name `constructor` is important here
  constructor(args, opts) {
    // Calling the super constructor is important so our generator is correctly set up
    super(args, opts);
  }

  initializing() {
    this.composeWith(require.resolve('../task-runner'));
  }

  async prompting() {
    var answers = await this.prompt([
      {
        type: "confirm",
        name: "genOpenAPI",
        message: "Enable REST API via OpenAPI generator?"
      }
    ]);
    this.templateConfig = { ...answers }

    if (!this.templateConfig.genOpenAPI) {
      return
    }

    answers = await this.prompt([
      {
        type: "input",
        name: "genOutputPath",
        default: "internal/web/rest",
        message: "Where do you want to place generated openapi go files?"
      },
      {
        type: "number",
        name: "restApiPort",
        default: 8080,
        message: "Which port will the api run?"
      }
    ])

    var openApiGenPackage = answers.genOutputPath.split("/").reverse()[0]
    this.templateConfig = {
      ...this.templateConfig,
      ...answers,
      openApiGenPackage,
      moduleName: sanitize.appName(this.appname)
    }

  }

  async configuring() {
  }


  async writing() {
    // generate only if needed
    if (!this.templateConfig.genOpenAPI) {
      return
    }

    await fs.readdir(this.templatePath("rest"), (err, files) => {
      if (err) {
        return err
      }
      files.map(filename => {
        console.log(filename)
        this.fs.copyTpl(this.templatePath(`rest/${filename}`),
          this.destinationPath(`${this.templateConfig.genOutputPath}/${filename}`),
          this.templateConfig)
      })
    })

    this.fs.copyTpl(this.templatePath(`config/config.go`),
      this.destinationPath(`cmd/config/${this.templateConfig.openApiGenPackage}.go`),
      this.templateConfig)

    this.fs.copyTpl(this.templatePath(`api/openapi.yaml`),
      this.destinationPath(`api/${this.templateConfig.openApiGenPackage}-openapi.yaml`),
      this.templateConfig)

    this.fs.copyTpl(this.templatePath(`openapi-gen.cfg.yaml`),
      this.destinationPath(`${this.templateConfig.openApiGenPackage}-openapi-gen.cfg.yaml`),
      this.templateConfig)

    this._updateMakefile()
  }

  install() {
    // install only if needed
    if (!this.templateConfig.genOpenAPI) {
      return
    }

    // make generate-
    let spinner = ora().start("Generating files from OpenAPI spec\n")
    this.spawnCommandSync("make", [`generate-${this.templateConfig.openApiGenPackage}`], { detached: false })
    spinner.succeed()
    // run go mod vendor
    spinner = ora().start("Running go mod vendor")
    this.spawnCommandSync("go", [`mod`, `vendor`])
    spinner.succeed()
  }
  end() {
  }

  async _updateMakefile() {
    var makefilePath = this.destinationPath('Makefile')

    if (this.fs.exists(makefilePath)) {
      // makefile exists, so append content to it
      var makefileContent = this.fs.read(makefilePath)
      const { ast } = parseMakefile(makefileContent)

      this._appendMakefileIfTargetDoesntExist(ast, makefilePath, `install-oapi-generator`, `makefile-install-oapi-gen.partial`)
      this._appendMakefileIfTargetDoesntExist(ast, makefilePath, `generate-${this.templateConfig.openApiGenPackage}`, `makefile-generate.partial`)
    } else {
      this.fs.copyTpl(this.templatePath('makefile-install-oapi-gen.partial'), this.destinationPath('Makefile'), this.templateConfig)
      this._appendMakefileTarget(makefilePath, `makefile-generate.partial`)
    }

  }

  _appendMakefileIfTargetDoesntExist(makefileAst, makefilePath, makefileTarget, makefilePartialTemplatePath) {
    var targetFound = makefileAst.find((entry) => {
      if (entry && entry.target == makefileTarget) {
        return true
      }
      false
    })
    if (!targetFound) {
      this._appendMakefileTarget(makefilePath, makefilePartialTemplatePath)
    }
  }

  _appendMakefileTarget(makefilePath, partialFilePath) {
    // create temporary makefile using templating
    this.fs.copyTpl(this.templatePath(partialFilePath), this.destinationPath(`.tmp/${partialFilePath}`), this.templateConfig)
    this.fs.append(makefilePath, this.fs.read(this.destinationPath(`.tmp/${partialFilePath}`)))
    this.fs.delete(this.destinationPath(`.tmp/${partialFilePath}`))
  }
};
