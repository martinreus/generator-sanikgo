var Generator = require('yeoman-generator');
const simpleGit = require('simple-git');

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
    new Promise((resolve, reject) => {
      var git = simpleGit()
      git.tag("", (err, data) => {
        if (err) {
          return reject(err);
        }
        return resolve(data)
      })
    }).then(tag => {
      this.log("git repository already initialised")
    }).catch(err => {
      this.log("Initialising git repository..")
      this.spawnCommandSync("git", [`init`])
      this.spawnCommandSync("git", [`add`, `.`])
      this.spawnCommandSync("git", [`commit`, `.`, `-m`, `Initial commit`])
      this.spawnCommandSync("git", [`tag`, "1.0.0"])
    })
  }

};
