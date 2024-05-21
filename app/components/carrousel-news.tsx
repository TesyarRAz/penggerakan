import React from "react";
import Slider from "react-slick";

interface CarouselNewsProps {
  items: {
    image: string;
    title: string;
    description: string;
  }[];
}

const settings = {
  dots: true,
  infinite: true,
  speed: 500,
  slidesToShow: 1,
  slidesToScroll: 1,
  autoplay: true,
  autoplaySpeed: 3000,
  arrows: true,
  responsive: [
    {
      breakpoint: 1024,
      settings: {
        slidesToShow: 2,
        slidesToScroll: 2,
        infinite: true,
        dots: true,
      },
    },
    {
      breakpoint: 600,
      settings: {
        slidesToShow: 1,
        slidesToScroll: 1,
        initialSlide: 1,
      },
    },
  ],
};

const CarouselNews: React.FC<CarouselNewsProps> = ({ items }) => {
  return (
    <div className="w-[700px] mx-auto pt-2 pb-10">
      <Slider {...settings}>
        {items.map((item, index) => (
          <div key={index} className="p-2">
            <div className="flex bg-indigo-600 rounded-xl h-full">
              <div className="w-52 h-44">
                <img
                  src={item.image}
                  alt={item.title}
                  className="rounded-l-lg w-full h-full object-fill"
                />
              </div>
              <div className="flex flex-col flex-1 my-2 mx-2 justify-between">
                <div>
                  <p className="text-2xl font-bold text-white">{item.title}</p>
                  <p className="font-medium text-white">{item.description}</p>
                </div>
                <div className="">
                  <button className="bg-white rounded-md p-1.5 text-xs">
                    Read More
                  </button>
                </div>
              </div>
            </div>
          </div>
        ))}
      </Slider>
    </div>
  );
};

export default CarouselNews;
