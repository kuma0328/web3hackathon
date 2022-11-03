import Head from 'next/head';
import { Page } from '../components/Wrapper/Page';
import { useGetHealth } from '../hooks/Health/useGetHealth';
import { Pagination, Navigation, Autoplay, EffectCube } from 'swiper';
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css/effect-cube';
import 'swiper/css';
import { LinkTo } from '../components/BaseParts/LinkTo';
export default function Home() {
  const { data, error, isLoading } = useGetHealth();
  return (
    <div>
      <Head>
        <title>fooDAO</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Page wide={true}>
        <main className=" bg-base-bg-color">
          <h1 className="text-center text-3xl font-extralight">FooDAO</h1>

          <div className="my-5 rounded-full bg-white p-20">
            <Swiper
              spaceBetween={30}
              centeredSlides={true}
              autoplay={{
                delay: 2500,
                disableOnInteraction: false,
              }}
              effect={'cube'}
              grabCursor={true}
              cubeEffect={{
                shadow: true,
                slideShadows: true,
                shadowOffset: 30,
                shadowScale: 5,
              }}
              pagination={{
                clickable: true,
              }}
              navigation={true}
              modules={[EffectCube, Autoplay]}
              className="h-52 w-52 rounded-md shadow-2xl"
            >
              <SwiperSlide>
                <img src="https://swiperjs.com/demos/images/nature-1.jpg" />
              </SwiperSlide>
              <SwiperSlide>
                <img src="https://swiperjs.com/demos/images/nature-2.jpg" />
              </SwiperSlide>
              <SwiperSlide>
                <img src="https://swiperjs.com/demos/images/nature-3.jpg" />
              </SwiperSlide>
              <SwiperSlide>
                <img src="https://swiperjs.com/demos/images/nature-4.jpg" />
              </SwiperSlide>
            </Swiper>
          </div>

          <LinkTo link="/community" className="text-center">
            始める
          </LinkTo>
        </main>
      </Page>
    </div>
  );
}
