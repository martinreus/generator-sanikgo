var Generator = require('yeoman-generator');

module.exports = class extends Generator {

  constructor(args, opts) {
    super(args, opts);
  }

  initializing() {
    this.composeWith(require.resolve('../base'));
  }

  async prompting() {
  }

  configuring() { }


  writing() { }

  install() {
  }

  end() {
    // run go mod vendor
    this.log("Initialising git repository..")
    this.spawnCommandSync("git", [`init`])
    this.spawnCommandSync("git", [`add`, `.`])
    this.spawnCommandSync("git", [`commit`, `.`, `-m`, `Initial commit`])
    this.spawnCommandSync("git", [`tag`, "1.0.0"])
  }

};
