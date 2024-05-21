"use client";

import React, { useEffect, useState } from "react";
import CarrouselNews from "../components/carrousel-news";
import { FaSearch } from "react-icons/fa";
import { RiEqualizerFill } from "react-icons/ri";
import CourseCard from "../components/card";

const carrouselItems = [
  {
    title: "First Slide",
    description: "this is the first slide",
    image: "/images/cat-1.jpg",
  },
  {
    title: "Second Slide",
    description: "this is the second slide",
    image: "/images/cat-2.jpg",
  },
  {
    title: "Third Slide",
    description: "this is the third slide",
    image: "/images/cat-3.jpg",
  },
];

const courseCardItems = [
  {
    title: "Course1",
    description: "this is course 1",
    image: "/images/cat-2.jpg",
  },
  {
    title: "Course2",
    description: "this is course 2",
    image: "/images/cat-2.jpg",
  },
  {
    title: "Course3",
    description: "this is course 3",
    image: "/images/cat-2.jpg",
  },
  {
    title: "Course4",
    description: "this is course 4",
    image: "/images/cat-2.jpg",
  },
  {
    title: "Course5",
    description: "this is course 5",
    image: "/images/cat-2.jpg",
  },
  {
    title: "Course6",
    description: "this is course 6",
    image: "/images/cat-2.jpg",
  },
];

const DashboardAdmin = () => {
  const [themeDark, setThemeDark] = useState(true);

  useEffect(() => {
    if (window.matchMedia("prefers-color-scheme:dark").matches) {
      setThemeDark(true);
    } else {
      setThemeDark(false);
    }
  }, []);

  useEffect(() => {
    if (themeDark) {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
  }, [themeDark]);

  return (
    <div className="bg-white dark:bg-gray-900 h-screen overflow-auto scrollbar-thin dark:scrollbar-thin-dark scroll-smooth">
      <div className="bg-gray-200 mt-5 mx-5 p-2 rounded-lg dark:bg-gray-800">
        <h2 className="font-sans font-bold text-3xl ml-5 dark:text-white">
          Home Page
        </h2>
      </div>

      <div className="justify-center items-center bg-gray-200 mx-5 rounded-lg mt-3 dark:bg-gray-800">
        {/* carrousel content */}
        <CarrouselNews items={carrouselItems} />
      </div>
      <div className="justify-between items-center mt-3">
        <div className="bg-gray-200 mt-5 mx-5 p-2 rounded-t-lg dark:bg-gray-800">
          <h2 className="font-sans font-bold text-3xl ml-5 dark:text-white">
            Browse Course
          </h2>
          <form action="#" className="flex mt-5">
            <input
              type="text"
              id="browseCourse"
              name="browseCourse"
              placeholder="Browse Course..."
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 ml-5 mr-1"
            />
            <div className="flex mr-4">
              <button
                type="button"
                className="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700 px-5 py-2.5 mr-1"
              >
                <FaSearch />
              </button>
              <button
                id="dropdownDefaultButton"
                data-dropdown-toggle="dropdown"
                className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm  text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 p-3 mr-1"
                type="button"
              >
                <RiEqualizerFill />
              </button>
            </div>

            {/* <!-- Dropdown menu --> */}
            <div
              id="dropdown"
              className="z-10 hidden bg-white divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700"
            >
              <ul
                className="py-2 text-sm text-gray-700 dark:text-gray-200"
                aria-labelledby="dropdownDefaultButton"
              >
                <li>
                  <a
                    href="#"
                    className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                  >
                    Math
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                  >
                    Science
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                  >
                    History
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                  >
                    Biology
                  </a>
                </li>
              </ul>
            </div>
          </form>
          {/* Card Course */}
          <div className="grid grid-cols-1 md:grid-cols-3 ml-2">
            {courseCardItems.map((cardItem, index) => (
              <CourseCard
                title={cardItem.title}
                description={cardItem.description}
                image={cardItem.image}
              />
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default DashboardAdmin;
