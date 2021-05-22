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
  }

  async prompting() {
    this.templateConfig = {}
    var answers = await this.prompt([
      {
        type: "input",
        name: "genOutputPath",
        default: "internal/restapi",
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
    var openApiGenPackageCamel = _.upperFirst(openApiGenPackage)
    var openApiGenPackageUpper = _.upperCase(openApiGenPackage)
    this.templateConfig = {
      ...this.templateConfig,
      ...answers,
      openApiGenPackage,
      openApiGenPackageCamel,
      openApiGenPackageUpper,
      moduleName: sanitize.appName(this.appname)
    }

  }

  async configuring() {
  }


  async writing() {

    this.log("openapi: writing files")
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

    this.log("openapi: writing files finished")
  }

  install() {
    this.log.write("openapi: installing...")
    // make generate-
    this.log.write("openapi: generating files from openapi definition...")
    this.spawnCommandSync("make", [`generate-${this.templateConfig.openApiGenPackage}`], { detached: false })
    this.log.write("openapi: openapi files generated")


    // run go mod vendor
    this.log("openapi: running go mod vendor")
    this.spawnCommandSync("go", [`mod`, `vendor`])
    this.log("openapi: go mod done")

  }
  end() {
  }

  async _updateMakefile() {
    var makefilePath = this.destinationPath('Makefile')

    if (this.fs.exists(makefilePath)) {
      // makefile exists, so append content to it
      this.log("openapi: makefile exists, adding targets..")
      this._appendMakefileIfTargetDoesntExist(
        makefilePath, `install-oapi-generator`, `makefile-install-oapi-gen.partial`, this.templateConfig)
      this._appendMakefileIfTargetDoesntExist(
        makefilePath, `generate-${this.templateConfig.openApiGenPackage}`, `makefile-generate.partial`, this.templateConfig)
    } else {
      this.log("openapi: makefile doesn't exist, creating new and generating targets.. ")
      this.fs.copyTpl(this.templatePath('makefile-install-oapi-gen.partial'), makefilePath, this.templateConfig)
      this._appendMakefileTarget(makefilePath, `makefile-generate.partial`, this.templateConfig)
    }

  }

};
