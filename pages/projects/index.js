import Image from "next/image";
import { useEffect, useState } from "react";
import Layout from "../../components/Layout";
import fs from "fs";
import matter from "gray-matter";
import Link from "next/link";
import { NextSeo } from "next-seo";

export default function Projects({ projects }) {
  const [search, setSearch] = useState("");
  const [selectedCategory, setSelectedCategory] = useState([]);
  const toggleClickCategoryCheckbox = (e) => {
    if (e.target.checked) {
      setSelectedCategory([...selectedCategory, parseInt(e.target.value)]);

      return;
    }

    setSelectedCategory(
      selectedCategory.filter((category) => {
        return category !== parseInt(e.target.value);
      })
    );
    return;
  };

  const computedProjects = projects
    .filter((project) => {
      if (selectedCategory.length === 0) return true;
      return project.categories.some((category) => {
        return selectedCategory.includes(category);
      });
    })
    .filter((project) => {
      if (search === "") return true;
      return (
        project.title.includes(search) ||
        (project.excerpt &&
          project.excerpt.toLowerCase().includes(search.toLowerCase())) ||
        (project.content &&
          project.content.toLowerCase().includes(search.toLowerCase()))
      );
    });

  return (
    <>
      <NextSeo
        title="Projects | SawirStudio"
        description="Projects we have worked on"
      />
      <Layout>
        <section id="banner" className="banner" role="banner">
          <div className="p-8">
            <h1 className="text-3xl font-bold">Latest projects</h1>
            <p className="text-xl">
              A collection of the latest projects from SawirStudio
            </p>
          </div>
        </section>
        <section>
          <div className="p-8 space-y-4">
            <input
              id="search"
              className="border border-gray-300 rounded-md px-4 py-2"
              placeholder="Search..."
              onInput={(e) => {
                setSearch(e.target.value);
              }}
              value={search}
            />
            <div className="flex space-x-8">
              <div className="space-x-2">
                <input
                  id="application"
                  type="checkbox"
                  value="3"
                  onChange={toggleClickCategoryCheckbox}
                  checked={selectedCategory.includes(3)}
                />
                <label htmlFor="application">APPLICATION</label>
              </div>
              <div className="space-x-2">
                <input
                  id="website"
                  type="checkbox"
                  value="2"
                  onChange={toggleClickCategoryCheckbox}
                  checked={selectedCategory.includes(2)}
                />
                <label htmlFor="website">WEBSITE</label>
              </div>
              <div className="space-x-2">
                <input
                  id="design"
                  type="checkbox"
                  value="1"
                  onChange={toggleClickCategoryCheckbox}
                  checked={selectedCategory.includes(1)}
                />
                <label htmlFor="design">DESIGN</label>
              </div>
            </div>
            {computedProjects.length === 0 && (
              <div className="text-center">
                <p>No projects found. Try to clear the filters...</p>
                <button
                  onClick={() => {
                    setSearch("");
                    setSelectedCategory([]);
                  }}
                >
                  Clear Filter
                </button>
              </div>
            )}
            {computedProjects.length > 0 && (
              <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                {computedProjects.map((project, index) => {
                  return (
                    <div key={`project_${index}`} className="cursor-pointer">
                      <Link href={`/projects/${project.slug}`}>
                        <div className="bg-white rounded shadow overflow-hidden">
                          <div className="relative">
                            <div className="pb-[56.25%]"></div>
                            <Image
                              className="w-full h-full object-cover"
                              alt={project.title}
                              src={project.featured_media}
                              layout="fill"
                            />
                          </div>
                          <div className="p-6">
                            <h3 className="text-xl font-bold">
                              {project.title}
                            </h3>
                            <p className="text-sm"></p>
                          </div>
                        </div>
                      </Link>
                    </div>
                  );
                })}
              </div>
            )}
          </div>
        </section>
      </Layout>
    </>
  );
}

export async function getStaticProps() {
  const projects = fs
    .readdirSync("pages/projects")
    .filter((filename) => {
      return filename.includes(".mdx");
    })
    .map((filename) => {
      const { data: frontmatter, content } = matter(
        fs.readFileSync(`pages/projects/${filename}`)
      );

      return {
        ...frontmatter,
        content,
        slug: filename.replace(/\.mdx/, ""),
      };
    });

  return {
    props: {
      projects: projects.sort((a, b) => {
        return new Date(b.published_at) - new Date(a.published_at);
      }),
    },
  };
}
