var Generator = require('yeoman-generator');

module.exports = class extends Generator {

  // The name `constructor` is important here
  constructor(args, opts) {
    // Calling the super constructor is important so our generator is correctly set up
    super(args, opts);

  }

  initializing() {

  }

  async prompting() {
    this.templateConfig = await this.prompt([
      {
        type: "input",
        name: "moduleName",
        message: "Module's name",
        default: this.appname.replace(" ", "-") // Default to current folder name
      }
    ]);

    this.log("app name", this.templateConfig.moduleName);
  }

  configuring() { }


  writing() { }

  install() { }
  end() {
    console.log("finished base app generation");
  }

};
