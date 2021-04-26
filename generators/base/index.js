var Generator = require('yeoman-generator');
var sanitize = require("../sanitize")

module.exports = class extends Generator {

  // The name `constructor` is important here
  constructor(args, opts) {
    // Calling the super constructor is important so our generator is correctly set up
    super(args, opts);

  }

  initializing() {
    this.templateConfig = {
      moduleName: sanitize.appName(this.appname) // Default to current folder name
    }
  }

  async prompting() {
  }

  configuring() { }


  writing() {
    this.fs.copyTpl(this.templatePath(`go.mod`),
      this.destinationPath(`go.mod`),
      this.templateConfig)

    await fs.readdir(this.templatePath("cmd"), (err, files) => {
      if (err) {
        return err
      }

      files.map(filename => {
        console.log(filename)
        // this.fs.copyTpl(this.templatePath(`cmd/${filename}`),
        //   this.destinationPath(`${this.templateConfig.genOutputPath}/${filename}`),
        //   this.templateConfig)
      })
    })

  }

  install() { }
  end() {
  }

  _copyFiles(from, to) {

  }
};
