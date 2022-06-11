import Link from "next/link";

export function Footer() {
  return (
    <footer className="bg-gray-50 gradient shadow-inner-light | print:shadow-none print:bg-transparent print:gradient-none">
      <div className="flex-none pt-16 pb-8 | print:pb-2" role="navigation">
        <div className="px-8 links links-gray text-gray leading-loose | md:leading-normal">
          <div className="grid grid-cols-2 items-start text-sm | md:flex md:justify-between">
            <div className="grid | md:grid-flow-col ml:items-center md:gap-12">
              <Link className="cursor-pointer" href="/about-us">
                About
              </Link>
              <Link className="cursor-pointer" href="/projects">
                Projects
              </Link>
              <a href="mailto:sawir.ricardo@gmail.com">Contact Us</a>
            </div>
          </div>
          <hr
            className="my-8 h-2px text-gray opacity-25 rounded | print:text-black"
            style={{ pageBreakAfter: "avoid" }}
          />
          <div className="grid gap-4 text-sm | sm:grid-cols-2 sm:gap-8 | md:flex flex-row-reverse justify-between">
            <address className="grid gap-4 | sm:gap-0 | md:grid-flow-col md:gap-8 md:text-right">
              <div></div>
            </address>
            <ul className="hidden | md:grid md:grid-flow-col md:gap-6 | print:hidden">
              <li>
                <a
                  href="https://github.com/sawirricardo"
                  target="_blank"
                  rel="nofollow noreferrer noopener"
                >
                  GitHub
                </a>
              </li>
              <li>
                <a
                  href="https://www.instagram.com/rsawir"
                  target="_blank"
                  rel="nofollow noreferrer noopener"
                >
                  Instagram
                </a>
              </li>
              <li>
                <a
                  href="https://twitter.com/RicardoSawir"
                  target="_blank"
                  rel="nofollow noreferrer noopener"
                >
                  Twitter
                </a>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </footer>
  );
}
