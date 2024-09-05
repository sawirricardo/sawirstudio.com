import Layout from "../components/Layout";
import db from "../database/db.json";
import fs from "fs";
import matter from "gray-matter";
import Link from "next/link";
import Image from "next/image";
import Cta from "../components/Cta";
import { NextSeo } from "next-seo";

export default function Home({ clients, projects }) {
  return (
    <div>Hello World!</div>
  )
  return (
    <>
      <NextSeo
        title="SawirStudio | Solid expertise in Application / Website / Digital Design"
        description="Crafting applications, websites, digital designs with heart."
      />
      <Layout>
        <section id="banner" className="banner" role="banner">
          <div className="p-8 space-y-8">
            <h1 className="text-4xl md:text-8xl">
              Solid expertise <br />
              in Application / Website / Digital Design
            </h1>
            <p className="text-xl md:text-2xl">
              Crafting applications, websites, digital designs with heart.
            </p>
          </div>
        </section>
        <section>
          <div className="p-8">
            <h2 className="text-xl font-bold">Some of our clients</h2>
            <div className="grid grid-cols-1 md:grid-cols-12 md:gap-4">
              {clients.map((clients, index) => {
                return (
                  <div key={`clients_${index}`} className="md:col-span-4">
                    <ul>
                      {clients.map((client, index) => {
                        return (
                          <li key={`client_${index}`}>
                            <h4 className="font-bold text-gray-800 text-sm">
                              {client.name}
                            </h4>
                            <p className="text-gray-600 text-xs"></p>
                          </li>
                        );
                      })}
                    </ul>
                  </div>
                );
              })}
            </div>
          </div>
        </section>
        <section id="portfolio">
          <div className="p-8 grid grid-cols-1 gap-8">
            {projects.map((project, index) => {
              return (
                <div
                  key={`project_${index}`}
                  className="grid grid-cols-1 md:grid-cols-8 gap-8 px-8"
                >
                  <div className="cursor-pointer sm:col-span-4 | md:col-start-2 md:col-span-3 md:-ml-32 md:mr-16">
                    <Link
                      href={`/projects/${project.slug}`}
                      className=" illustration is-left h-full"
                      target="_blank"
                      rel="nofollow noreferrer noopener"
                    >
                      <div className="relative shadow-xl">
                        <div className="pb-[56.25%]"></div>
                        <Image
                          className="w-full h-full object-cover"
                          alt={project.tilte}
                          src={project.featured_media}
                          layout="fill"
                        />
                      </div>
                    </Link>
                  </div>
                  <div className="cursor-pointer sm:col-span-3 sm:col-start-5 | self-center">
                    <Link href={`/projects/${project.slug}`}>
                      <div className="cursor-pointer markup links-blue links-underline | sm:grid-text-right">
                        <h3 className="text-xl md:text-3xl font-bold">
                          {project.title}
                        </h3>
                        <p className="text-lg">{project.excerpt}</p>
                      </div>
                    </Link>
                  </div>
                </div>
              );
            })}
          </div>
          <div className="text-center text-2xl text-blue-500">
            <Link className="cursor-pointer" href="/projects">
              See all projects
            </Link>
          </div>
        </section>
        <Cta />
      </Layout>
    </>
  );
}

export async function getStaticProps() {
  const projects = fs
    .readdirSync("pages/projects")
    .filter((filename) => {
      return filename.includes(".md");
    })
    .map((filename) => {
      const { data: frontmatter } = matter(
        fs.readFileSync(`pages/projects/${filename}`)
      );

      return {
        ...frontmatter,
        slug: filename.replace(/\.mdx/, ""),
      };
    });

  const chunk = (input, size) => {
    return input.reduce((arr, item, idx) => {
      return idx % size === 0
        ? [...arr, [item]]
        : [...arr.slice(0, -1), [...arr.slice(-1)[0], item]];
    }, []);
  };
  return {
    props: {
      clients: chunk(db.clients, 4),
      projects: projects
        .sort((a, b) => {
          return new Date(b.published_at) - new Date(a.published_at);
        })
        .slice(0, 5),
    },
  };
}
