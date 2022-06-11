export default function Cta() {
  return (
    <section id="cta">
      <div className="p-8">
        <div className="bg-gradient-to-r from-green-600 to-green-400 text-white rounded-lg shadow-xl">
          <div className="p-8 grid md:grid-cols-2 md:items-end">
            <div className="links-underline links-white">
              <p className="text-xl font-bold">Enough about us.</p>
            </div>
            <h2 className="md:ml-auto text-2xl md:text-6xl md:text-right">
              <a
                className="hover:text-gray-200 underline"
                href="mailto:sawir.ricardo@gmail.com"
              >
                How can we
                <br />
                help you
              </a>
              ?
            </h2>
          </div>
        </div>
      </div>
    </section>
  );
}
