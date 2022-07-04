import { NextSeo } from "next-seo";
import Link from "next/link";
import fs from "fs";
import Layout from "../../components/Layout";
import matter from "gray-matter";

export default function Posts({ posts }) {
  return (
    <>
      <NextSeo
        title="Blog | SawirStudio"
        description="Crafting applications, websites, digital designs with heart."
      />
      <Layout>
        <section>
          <div className="p-8">
            <h1 className="text-3xl font-bold text-center mb-8">Blog</h1>

            <ul className="max-w-2xl mx-auto text-xl">
              {posts.map((post, index) => {
                return (
                  <li
                    key={`post_${index}`}
                    className="text-blue-500 hover:text-blue-400"
                  >
                    <Link href={`/posts/${post.slug}`}>
                      <a>
                        <span className="text-gray-400">
                          {post.published_at} -
                        </span>{" "}
                        {post.title}
                      </a>
                    </Link>
                  </li>
                );
              })}
            </ul>
          </div>
        </section>
      </Layout>
    </>
  );
}

export async function getStaticProps() {
  const posts = fs
    .readdirSync("pages/posts")
    .filter((filename) => {
      return filename.includes(".mdx");
    })
    .map((filename) => {
      const { data: frontmatter, content } = matter(
        fs.readFileSync(`pages/posts/${filename}`)
      );

      return {
        ...frontmatter,
        content,
        slug: filename.replace(/\.mdx/, ""),
      };
    });

  return {
    props: {
      posts: posts.sort((a, b) => {
        return new Date(b.published_at) - new Date(a.published_at);
      }),
    },
  };
}
