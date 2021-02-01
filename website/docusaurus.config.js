module.exports = {
  title: "LeetCode Go",
  tagline: "LeetCode golang Solutions",
  url: "https://leetcode.xulingming.cn",
  baseUrl: "/",
  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",
  favicon: "img/favicon.ico",
  organizationName: "facebook", // Usually your GitHub org/user name.
  projectName: "docusaurus", // Usually your repo name.
  themeConfig: {
    prism: {
      theme: require("./src/js/monokaiTheme.js"),
    },
    navbar: {
      title: "LeetCode Go",
      logo: {
        alt: "LeetCode Logo",
        src: "img/favicon.ico",
      },
      items: [
        {
          to: "Algorithms/",
          activeBasePath: "Algorithms",
          label: "算法",
          position: "left",
        },
        {
          to: "Concurrency/",
          activeBasePath: "Concurrency",
          label: "并发",
          position: "left",
        },
        {
          to: "Database/",
          activeBasePath: "Database",
          label: "数据库",
          position: "left",
        },
        {
          to: "Shell/",
          activeBasePath: "Shell",
          label: "Shell 脚本",
          position: "left",
        },
        {
          to: "LCCI/",
          activeBasePath: "LCCI",
          label: "面试题",
          position: "left",
        },
        {
          to: "LCOF/",
          activeBasePath: "LCOF",
          label: "剑指 Offer",
          position: "left",
        },
        {
          href: "https://github.com/jackdon/leetcode",
          label: "Github",
          position: "right",
        },
      ],
    },
    footer: {
      links: [
        {
          title: "LeetCode",
          items: [
            {
              label: "算法",
              to: "Algorithms/",
            },
            {
              label: "并发",
              to: "Concurrency/",
            },
            {
              label: "数据库",
              to: "Database/",
            },
            {
              label: "Shell 脚本",
              to: "Shell/",
            },
            {
              label: "剑指 Offer",
              to: "LCCI/",
            },
          ],
        },
        {
          title: "更多",
          items: [
            {
              label: "更新日志",
              to: "blog",
            },
            {
              label: "GitHub",
              href: "https://github.com/jackdon/leetcode",
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} LeetCode Go.`,
    },
    algolia: {
      apiKey: '99b342d9ee4c048c72a0e028e54201a9',
      appId: 'PIBA4C4WI9',
      indexName: 'leetcode',
      // Optional: see doc section bellow
      algoliaOptions: {}
    }  
  },
  presets: [
    [
      "@docusaurus/preset-classic",
      {
        docs: {
          path: "./docs",
          routeBasePath: "/",
          sidebarPath: require.resolve("./sidebars.js"),
          // Please change this to your repo.
          editUrl:
            "https://github.com/facebook/docusaurus/edit/master/website/",
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          editUrl:
            "https://github.com/facebook/docusaurus/edit/master/website/blog/",
        },
        theme: {
          customCss: require.resolve("./src/css/custom.css"),
        },
      },
    ],
  ],
  themes: ["@docusaurus/theme-live-codeblock"],
};
