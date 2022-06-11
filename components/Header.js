import Link from "next/link";

export function Header() {
  return (
    <header className="pt-8 flex-none z-10 | md:bg-white md:shadow-light md:py-8 | print:bg-transparent print:shadow-none">
      <div className="px-8 leading-loose | md:leading-none md:flex md:items-stretch">
        <Link
          className="cursor-pointer flex-shrink-0 logo h-8 w-20 mr-16 mb-8 block | md:mb-0 md:w-48 md:h-auto"
          href="/"
          title="Home"
        >
          <span className="text-3xl md:text-4xl lg:text-5xl xl:text-6xl font-bold">
            SawirStudio
          </span>
        </Link>
        <div className="grid grid-cols-2 items-start col-gap-8 | md:grid-cols-1 md:row-gap-6 md:justify-end md:justify-items-end md:ml-auto">
          <nav className="flex links links-black | md:row-start-2">
            <ul className=" md:grid md:grid-flow-col md:gap-12 md:justify-between | print:hidden">
              <li>
                <Link className="cursor-pointer" href="/">
                  Home
                </Link>
              </li>
              <li>
                <Link className="cursor-pointer" href="/projects">
                  Projects
                </Link>
              </li>
              <li>
                <Link className="cursor-pointer" href="/about-us">
                  About us
                </Link>
              </li>
              <li>
                <a href="mailto:sawir.ricardo@gmail.com">Contact Us</a>
              </li>
            </ul>
          </nav>
        </div>
      </div>
      <div className="wrap | md:hidden | print:hidden">
        <hr className="mt-8 h-2px opacity-25 rounded text-gray" />
      </div>
    </header>
  );
}
