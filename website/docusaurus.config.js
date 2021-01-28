module.exports = {
  title: 'LeetCode Go',
  tagline: 'LeetCode golang Solutions',
  url: 'https://leetcode.xulingming.cn',
  baseUrl: '/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'facebook', // Usually your GitHub org/user name.
  projectName: 'docusaurus', // Usually your repo name.
  themeConfig: {
    prism: {
      theme: require('./src/js/monokaiTheme.js')
    },
    navbar: {
      title: 'LeetCode Go',
      logo: {
        alt: 'LeetCode Logo',
        src: 'img/logo.svg',
      },
      items: [
        {
          to: 'Algorithms/',
          activeBasePath: 'Algorithms',
          label: '算法',
          position: 'left',
        },
        {
          to: 'LCCI/',
          activeBasePath: 'LCCI',
          label: '面试题',
          position: 'left',
        },
        {
          to: 'LCOF/',
          activeBasePath: 'LCOF',
          label: '剑指 Offer',
          position: 'left',
        },
        {
          to: 'Concurrency/',
          activeBasePath: 'Concurrency',
          label: '并发',
          position: 'left',
        },
        {
          to: 'Database/',
          activeBasePath: 'Database',
          label: '数据库',
          position: 'left',
        },
        {
          to: 'Shell/',
          activeBasePath: 'Shell',
          label: 'Shell 脚本',
          position: 'left',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'LeetCode',
          items: [
            {
              label: '算法',
              to: 'Algorithms/',
            },
            {
              label: 'Shell 脚本',
              to: 'Shell/',
            },
            {
              label: '并发',
              to: 'Concurrency/',
            },
            {
              label: '数据库',
              to: 'Database/',
            },
          ],
        },
        {
          title: '更多',
          items: [
            {
              label: '更新日志',
              to: 'blog',
            },
            {
              label: 'GitHub',
              href: 'https://github.com/jackdonw/leetcode-2011',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} LeetCode Go.`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          path: './docs',
          routeBasePath: '/',
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl:
            'https://github.com/facebook/docusaurus/edit/master/website/',
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          editUrl:
            'https://github.com/facebook/docusaurus/edit/master/website/blog/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
  themes: ['@docusaurus/theme-live-codeblock']
};
