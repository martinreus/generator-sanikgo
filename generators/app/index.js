var Generator = require('yeoman-generator');
let ora = require("ora");

module.exports = class extends Generator {

  constructor(args, opts) {
    super(args, opts);

  }

  initializing() {
    this.composeWith(require.resolve('../base'));
    this.composeWith(require.resolve('../task-runner'));
    this.composeWith(require.resolve('../openapi'));
  }

  async prompting() {
  }

  configuring() { }


  writing() { }

  install() {
    // run go mod vendor
    let spinner = ora().start("Running go mod vendor")
    this.spawnCommandSync("go", [`mod`, `vendor`])
    spinner.succeed()
  }
  end() {
  }

};
