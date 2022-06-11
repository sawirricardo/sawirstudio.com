import Image from "next/image";
import Cta from "../../components/Cta";
import Layout from "../../components/Layout";
import md from "markdown-it";
import matter from "gray-matter";
import fs from "fs";
import db from "../../database/db.json";

export default function ViewProject({ frontmatter, content, client }) {
  return (
    <>
      <Layout>
        <section id="banner" className="banner" role="banner">
          <div className="p-8 max-w-3xl">
            <h1 className="text-3xl font-extrabold mb-3">
              {frontmatter.title}
            </h1>
            <p className="text-xl">{frontmatter.excerpt}</p>
          </div>
        </section>
        <section>
          <div className="p-8">
            <div className="grid grid-cols-1 md:grid-cols-12 gap-8">
              <div className="hidden md:block md:col-span-2"></div>
              <div
                className="md:max-w-xl prose lg:mx-auto col-span-1 md:col-span-8"
                dangerouslySetInnerHTML={{ __html: md().render(content) }}
              ></div>
              <div className="col-span-1 md:col-span-2">
                <ul>
                  {client && (
                    <>
                      <li className="break-all">
                        <strong>Client</strong>: {client.name}
                      </li>
                    </>
                  )}
                  {frontmatter.website_url && (
                    <li className="break-all">
                      <strong>Website URL</strong>:
                      <a
                        href={frontmatter.website_url}
                        target="_blank"
                        rel="noreferrer"
                      >
                        {frontmatter.website_url}
                      </a>
                    </li>
                  )}
                  {frontmatter.partner && (
                    <li className="break-all">
                      <strong>Partner</strong>:{frontmatter.partner}
                    </li>
                  )}
                  {frontmatter.google_playstore_url && (
                    <li className="break-all">
                      <strong>Google Playstore URL</strong>:
                      {frontmatter.google_playstore_url}
                    </li>
                  )}
                </ul>
              </div>
            </div>
          </div>
        </section>
        <Cta />
      </Layout>
    </>
  );
}

export async function getStaticPaths() {
  return {
    fallback: true, // false or 'blocking'
    paths: fs.readdirSync("pages/projects").map((fileName) => {
      return {
        params: {
          id: fileName.replace(".mdx", ""),
        },
      };
    }),
  };
}

export async function getStaticProps({ params: { id } }) {
  const fileName = fs.readFileSync(`pages/projects/${id}.mdx`, "utf-8");
  const { data: frontmatter, content } = matter(fileName);

  return {
    props: {
      frontmatter,
      content,
      client: db.clients.find((client) => client.id === frontmatter.client_id),
    },
  };
}
