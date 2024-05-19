import { useState, useEffect } from "react";
import { FaAngleLeft, FaAngleRight } from "react-icons/fa";

interface CarrouselProps {
  items: {
    title: string;
    description: string;
    image: string;
  }[];
}

function CarrouselNews({ items }: CarrouselProps) {
  const [currentIndex, setCurrentIndex] = useState(0);

  const handlePrev = () => {
    const newIndex = (currentIndex - 1 + items.length) % items.length;
    setCurrentIndex(newIndex);
  };

  const handleNext = () => {
    const newIndex = (currentIndex + 1) % items.length;
    setCurrentIndex(newIndex);
  };

  useEffect(() => {
    const interval = setInterval(() => {
      handleNext();
    }, 5000); // 5000ms = 5 seconds

    return () => {
      clearInterval(interval);
    };
  }, [currentIndex]);

  useEffect(() => {
    console.log("Current Index:", currentIndex);
    console.log("Current Item:", items[currentIndex]);
  }, [currentIndex, items]);

  return (
    <div className="relative w-full max-w-lg mx-auto">
      <div className="overflow-hidden relative rounded-lg shadow-lg">
        {items.map((item, index) => (
          <div
            key={index}
            className={`absolute inset-0 transition-transform duration-700 ease-in-out transform ${
              index === currentIndex
                ? "translate-x-0"
                : index < currentIndex
                ? "-translate-x-full"
                : "translate-x-full"
            }`}
            style={{
              zIndex: index === currentIndex ? 1 : 0,
            }}
          >
            <img
              src={item.image}
              alt={item.title}
              className="w-full h-64 object-cover"
            />
            <div className="absolute bottom-0 left-0 right-0 bg-black bg-opacity-50 text-white p-4">
              <h3 className="text-xl font-bold">{item.title}</h3>
              <p>{item.description}</p>
            </div>
          </div>
        ))}
      </div>
      <button
        onClick={handlePrev}
        className="absolute top-1/2 left-0 transform -translate-y-1/2 bg-white bg-opacity-75 p-2 rounded-full shadow-md"
      >
        <FaAngleLeft />
      </button>
      <button
        onClick={handleNext}
        className="absolute top-1/2 right-0 transform -translate-y-1/2 bg-white bg-opacity-75 p-2 rounded-full shadow-md"
      >
        <FaAngleRight />
      </button>
    </div>
  );
}

export default CarrouselNews;
