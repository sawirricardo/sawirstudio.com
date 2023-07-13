import { NextSeo } from "next-seo";
import Layout from "../components/Layout";

export default function AboutUs() {
  return (
    <>
      <NextSeo title="About | SawirStudio" description="About SawirStudio" />
      <Layout>
        <section id="banner" className="banner" role="banner">
          <div className="p-8 space-y-8">
            <h1 className="text-4xl md:text-8xl">Glad you found us</h1>
            <p className="text-xl md:text-2xl">
              SawirStudio helps you to achieve your business goals with
              applications, websites, and designs.
            </p>
          </div>
        </section>
        <section>
          <div className="p-8 space-y-8">
            <div className="grid grid-cols-1 md:grid-cols-12 gap-8">
              <div className="hidden md:col-span-2"></div>
              <div className="prose mx-auto md:col-span-8">
                <p>
                  We love to work in partnership: as advisors, as architects of
                  technically challenging projects.
                </p>
                <p>
                  We&apos;d really like to know what kind of project or problem
                  you&apos;re dealing with. Feed us with as much input as you
                  have, so we can get accurate early on.
                </p>
                <h3>Our technology stacks</h3>
                <h4>BACKEND</h4>
                <ul>
                  <li>PHP</li>
                  <li>Laravel</li>
                  <li>Golang</li>
                  <li>Javascript (NodeJS)</li>
                  <li>Ruby on Rails</li>
                  <li>Django</li>
                  <li>Python</li>
                  <li>MySQL</li>
                  <li>PostgreSQL</li>
                  <li>MongoDB</li>
                  <li>WordPress</li>
                </ul>
                <h4>FRONTEND</h4>
                <ul>
                  <li>Flutter</li>
                  <li>Livewire</li>
                  <li>AlpineJS</li>
                  <li>Vue.js</li>
                  <li>React.js</li>
                  <li>TailwindCSS</li>
                </ul>
              </div>
            </div>
          </div>
        </section>
      </Layout>
    </>
  );
}
