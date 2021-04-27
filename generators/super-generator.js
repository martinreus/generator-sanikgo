const Generator = require("yeoman-generator");
var fs = require('fs')
const parseMakefile = require('@kba/makefile-parser')

module.exports = class SuperGenerator extends Generator {
  // The name `constructor` is important here
  constructor(args, opts) {
    // Calling the super constructor is important so our generator is correctly set up
    super(args, opts);

  }

  async _copyFiles(from, to, templateConfig) {
    await fs.readdir(this.templatePath(from), (err, files) => {
      if (err) {
        return err
      }
      files.map(filename => {
        if (fs.lstatSync(`${from}/${filename}`).isDirectory()) {
          this._copyFiles(`${from}/${filename}`, `${to}/${filename}`)
        } else {
          this.fs.copyTpl(this.templatePath(`${from}/${filename}`),
            this.destinationPath(`${to}/${filename}`),
            templateConfig)
        }
      })
    })
  }

  // appends a makefilePartialTemplatePath to a makefilePath, under the makefileTarget label
  _appendMakefileIfTargetDoesntExist(makefilePath, makefileTarget, makefilePartialTemplatePath) {
    var makefileContent = this.fs.read(makefilePath)
    const { ast } = parseMakefile(makefileContent)

    var targetFound = ast.find((entry) => {
      if (entry && entry.target == makefileTarget) {
        return true
      }
      false
    })
    if (!targetFound) {
      this._appendMakefileTarget(makefilePath, makefilePartialTemplatePath)
    }
  }

  _appendMakefileTarget(makefilePath, partialFilePath) {
    // create temporary makefile using templating
    this.fs.copyTpl(this.templatePath(partialFilePath), this.destinationPath(`.tmp/${partialFilePath}`), this.templateConfig)
    this.fs.append(makefilePath, this.fs.read(this.destinationPath(`.tmp/${partialFilePath}`)))
    this.fs.delete(this.destinationPath(`.tmp/${partialFilePath}`))
  }
}