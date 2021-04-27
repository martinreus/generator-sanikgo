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
    this.log("base: writing files")
    await this._copyFiles(this.templatePath(""), this.destinationPath(""), this.templateConfig)
    this.log("base: write finished")
  }

  install() {
    // run go mod vendor
    this.log("base: running go mod vendor")
    this.spawnCommandSync("go", [`mod`, `vendor`])
    this.log("base: go mod done")
  }

  end() {
  }

};
