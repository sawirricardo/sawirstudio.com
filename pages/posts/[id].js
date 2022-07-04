import { NextSeo } from "next-seo";
import fs from "fs";
import Layout from "../../components/Layout";
import matter from "gray-matter";
import md from "markdown-it";
import Script from "next/script";
import ConvertKitForm from "../../components/ConvertKitForm";

export default function ShowPost({ frontmatter, content }) {
  return (
    <>
      <NextSeo
        title={frontmatter && frontmatter.title}
        description={
          (frontmatter && frontmatter.excerpt) ??
          `Crafting applications, websites, digital designs with heart.`
        }
        titleTemplate={`%s | SawirStudio`}
        defaultTitle="SawirStudio"
      />
      <Layout>
        <section>
          <div className="p-8">
            <div className="max-w-3xl  mx-auto">
              <h1 className="text-3xl font-extrabold mb-3 text-center">
                {frontmatter && frontmatter.title}
              </h1>
              <div
                className="prose prose-xl"
                dangerouslySetInnerHTML={{ __html: md().render(content ?? "") }}
              ></div>
            </div>
          </div>
        </section>
        <section className="mx-auto max-w-3xl p-8">
          <ConvertKitForm />
        </section>
      </Layout>
    </>
  );
}

export async function getStaticPaths() {
  return {
    fallback: true, // false or 'blocking'
    paths: fs
      .readdirSync("pages/posts")
      .filter((fileName) => {
        return fileName.includes(".mdx");
      })
      .map((fileName) => {
        return {
          params: {
            id: fileName.replace(".mdx", ""),
          },
        };
      }),
  };
}

export async function getStaticProps({ params: { id } }) {
  const fileName = fs.readFileSync(`pages/posts/${id}.mdx`);
  const { data: frontmatter, content } = matter(fileName);

  return {
    props: {
      frontmatter,
      content,
    },
  };
}
