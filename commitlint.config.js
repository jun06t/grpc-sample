module.exports = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    // 日本語の subject を許容するため case ルールを無効化
    'subject-case': [0],
    // 生成 AI による長文 body / footer を許容するため行長制限を無効化
    'body-max-line-length': [0],
    'footer-max-line-length': [0],
  },
};
