var Generator = require('yeoman-generator');

module.exports = class extends Generator {

  constructor(args, opts) {
    super(args, opts);
  }

  initializing() {
    return this.composeWith(require.resolve('../base')).composeWith(require.resolve('../openapi'));
  }

  async prompting() {
  }

  configuring() { }


  writing() { }

  install() {
  }

  end() {
  }

};
