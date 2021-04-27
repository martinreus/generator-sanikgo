var fs = require('fs')
var sanitize = require("../sanitize")
var _ = require('lodash');
const SuperGenerator = require('../super-generator');

module.exports = class extends SuperGenerator {

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
        default: "internal/web/restapi",
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
    var openApiGenPackageUpper = _.upperFirst(openApiGenPackage)
    this.templateConfig = {
      ...this.templateConfig,
      ...answers,
      openApiGenPackage,
      openApiGenPackageUpper,
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

    await this._copyFiles(this.templatePath("restapi/server"),
      this.destinationPath(`${this.templateConfig.genOutputPath}`), this.templateConfig)

    this.fs.copyTpl(this.templatePath(`restapi/instance/restapi.go`),
      this.destinationPath(`cmd/app/${this.templateConfig.openApiGenPackage}.go`),
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
    this.log.write("Generating files from openapi definition...")
    this.spawnCommandSync("make", [`generate-${this.templateConfig.openApiGenPackage}`], { detached: false })
  }
  end() {
  }

  async _updateMakefile() {
    var makefilePath = this.destinationPath('Makefile')

    if (this.fs.exists(makefilePath)) {
      // makefile exists, so append content to it
      this._appendMakefileIfTargetDoesntExist(makefilePath, `install-oapi-generator`, `makefile-install-oapi-gen.partial`)
      this._appendMakefileIfTargetDoesntExist(makefilePath, `generate-${this.templateConfig.openApiGenPackage}`, `makefile-generate.partial`)
    } else {
      this.fs.copyTpl(this.templatePath('makefile-install-oapi-gen.partial'), this.destinationPath('Makefile'), this.templateConfig)
      this._appendMakefileTarget(makefilePath, `makefile-generate.partial`)
    }

  }

};
