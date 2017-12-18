module.exports = {
  root: true,
  parser: 'babel-eslint',
  env: {
    browser: true,
    node: true
  },
  extends: 'standard',
  // required to lint *.vue files
  plugins: [
    'html'
  ],
  // add your custom rules here
  rules: {
      'indent': [2, 4, { 'SwitchCase': 2 }], // 缩进空格数
      'block-spacing': [2, 'always'], // 强制在单行代码块中使用一致的空格
      'arrow-spacing': [2, { 'before': true, 'after': true }], // 强制箭头函数的箭头前后使用一致的空格
      'comma-style': [2, 'last'], // 强制使用一致的逗号风格
      'eqeqeq': [2, 'allow-null'], // 要求使用 === 和 !==
      'quotes': [2, 'single', { 'avoidEscape': true, 'allowTemplateLiterals': true }], // 强制使用一致的反勾号、双引号或单引号
  },
  globals: {}
}
