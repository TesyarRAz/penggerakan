"use client";

import React, { useEffect, useState } from "react";
import Carrousel from "./carrousel";
import { FaSearch } from "react-icons/fa";
import { RiEqualizerFill } from "react-icons/ri";
import CourseCard from "../courses/_components/course-card";

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

const Dashboard = () => {
  const [ setTheme, theme ] = useState(true);

  return (
    <div className="bg-white dark:bg-gray-900">
      <div>
        <h2 className="font-sans font-bold text-3xl ml-5 pt-5 dark:text-white">
          Home Page
        </h2>
      </div>
      <div className="flex-1 pt-16 h-40 mt-5 mx-40 rounded-lg bg-gray-200 mb-5">
        {/* carrousel conten */}
        <Carrousel items={carrouselItems}></Carrousel>
      </div>
      <hr />
      
    </div>
  );
};

export default Dashboard;
