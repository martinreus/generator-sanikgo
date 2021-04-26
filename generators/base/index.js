var sanitize = require("../sanitize");
const SuperGenerator = require('../super-generator');


module.exports = class extends SuperGenerator {

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


  async writing() {
    this.fs.copyTpl(this.templatePath(`go.mod`),
      this.destinationPath(`go.mod`),
      this.templateConfig)

    await this._copyFiles(this.templatePath("cmd"), this.destinationPath("cmd"))
  }

  install() { }

  end() {
  }

};
