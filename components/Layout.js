import { NextSeo } from "next-seo";
import Head from "next/head";
import Script from "next/script";
import { Footer } from "./Footer";
import { Header } from "./Header";

export default function Layout({ children }) {
  return (
    <>
      <Script
        src="https://www.googletagmanager.com/gtag/js?id=UA-135085180-1"
        strategy="afterInteractive"
      />
      <Script id="google-analytics" strategy="afterInteractive">
        {`
          window.dataLayer = window.dataLayer || [];
          function gtag(){window.dataLayer.push(arguments);}
          gtag('js', new Date());

          gtag('config', 'UA-135085180-1');
        `}
      </Script>
      <Header />
      <main>{children}</main>
      <Footer />
    </>
  );
}
