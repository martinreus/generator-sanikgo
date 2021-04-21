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
    }

  }

  async prompting() {
  }

  configuring() { }


  writing() {

    await fs.readdir(this.templatePath(""), (err, files) => {
      if (err) {
        return err
      }
      files.map(filename => {
        console.log(filename)
        this.fs.copyTpl(this.templatePath(`${filename}`),
          this.destinationPath(`pkg/tasks/${filename}`),
          this.templateConfig)
      })
    })
  }

  install() { }
  end() {
    console.log("finished task runner generation");
  }

};
