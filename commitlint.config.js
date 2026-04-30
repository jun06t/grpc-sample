module.exports = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    // 日本語の subject を許容するため case ルールを無効化
    'subject-case': [0],
  },
};
