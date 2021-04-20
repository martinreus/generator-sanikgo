var Generator = require('yeoman-generator');

module.exports = class extends Generator {

  constructor(args, opts) {
    super(args, opts);

  }
  
  initializing() {
    this.composeWith(require.resolve('../base'));
    this.composeWith(require.resolve('../openapi'));
  }

  async prompting(){
  }

  configuring(){}

  
  writing(){}
  
  install(){}
  end(){
  }

};
