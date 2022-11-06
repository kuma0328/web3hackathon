import Head from 'next/head';
import { Page } from '../components/Wrapper/Page';
import { Autoplay } from 'swiper';
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css/effect-cube';
import 'swiper/css';
import { LinkTo } from '../components/BaseParts/LinkTo';
import { SwipeImg } from '../components/Home/SwipeImg';
import { TypoGraphy } from '../components/BaseParts/TypoGraphy';
import { getWindowSize } from '../hooks/useGetwidth';

export default function Home() {
  return (
    <div>
      <Head>
        <title>fooDAO</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Page wide={true}>
        <main className=" bg-base-bg-color">
          <h1 className="text-center text-3xl font-medium">FooDAO</h1>
          <TypoGraphy className="my-3 text-center font-thin">
            あなたと郷土料理をつなげるweb3コミュニティ
          </TypoGraphy>
          <div className="w-screen px-5 py-10 md:px-40">
            <Swiper
              slidesPerView={getWindowSize().width <= 960 ? 3 : 6}
              spaceBetween={30}
              autoplay={{
                delay: 2500,
                disableOnInteraction: false,
              }}
              loop={true}
              modules={[Autoplay]}
            >
              {[1, 2, 3, 4, 5, 6, 7].map((i) => {
                return (
                  <SwiperSlide key={i} className="my-auto">
                    <SwipeImg image={i} />
                  </SwiperSlide>
                );
              })}
            </Swiper>
          </div>

          <LinkTo link="/community" className="m-auto w-40 text-center">
            始める
          </LinkTo>
        </main>
      </Page>
    </div>
  );
}
