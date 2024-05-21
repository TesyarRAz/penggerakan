"use client";

import React, { useEffect, useState } from "react";
import { FaSearch } from "react-icons/fa";
import { RiEqualizerFill } from "react-icons/ri";
import CourseCard from "../courses/_components/course-card";
import Carousel from "./carousel";

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
const Dashboard = () => {
  const [ setTheme, theme ] = useState(true);

  return (
    <div className="bg-white dark:bg-gray-900">
      <div>
        <h2 className="font-sans font-bold text-3xl ml-5 pt-5 dark:text-white">
          Home Page
        </h2>
      </div>
      <div className="relative justify-center items-center bg-gray-200 mx-5 rounded-lg mt-3 dark:bg-gray-800">
        {/* carrousel content */}
        <Carousel items={carrouselItems} />
      </div>
      <hr />
      
    </div>
  );
};

export default Dashboard;
