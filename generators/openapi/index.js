var Generator = require('yeoman-generator');
var fs = require('fs')
var sanitize = require("../sanitize")
const parseMakefile = require('@kba/makefile-parser')


module.exports = class extends Generator {

  // The name `constructor` is important here
  constructor(args, opts) {
    // Calling the super constructor is important so our generator is correctly set up
    super(args, opts);

  }

  initializing() { }

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

    this.log("configuration chosen:", this.templateConfig);
  }

  async configuring() {
    await fs.stat(this.destinationPath(""), (err, stats) => {
      if (!err) {
        console.log("makefile does not exist, generating a new one");
        this.makefileExists = false
      } else {
        this.makefileExists = true
      }
    });
    this.makefileExists
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
  }

  install() {
    // install only if needed
    if (!this.templateConfig.genOpenAPI) {
      return
    }



    // go mod vendor
  }
  end() {
    this.log("finished open api");
  }


};
