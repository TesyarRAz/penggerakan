import React from "react";
import CourseCard from "./course-card";
import { FaPlus, FaSearch } from "react-icons/fa";
import { RiEqualizerFill } from "react-icons/ri";
import Link from "next/link";

const CourseList = ({
    courses,
}: {
    courses: CourseResponse[]
}) => {
  return (
    <div className="flex-auto flex-col">
      <div className="flex justify-between items-center">
        <h2 className="font-sans font-bold text-3xl ml-5 dark:text-white">
          Browse Course
        </h2>
        <div>
            <Link href="/dashboard/courses/create" className="flex items-center gap-4 bg-gray-800 py-2 px-3 rounded-sm hover:bg-gray-700">
                <FaPlus />
                <span>Tambah</span>
            </Link>
        </div>
      </div>
      <div className="justify-between items-center mt-3">
        <form action="#" className="flex">
          <input
            type="text"
            id="browseCourse"
            name="browseCourse"
            placeholder="Browse Course..."
            className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-full focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 ml-5 mr-1"
          />
          <button
            type="button"
            className="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-full text-sm dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700 p-3 mr-1"
          >
            <FaSearch />
          </button>
          <button
            id="dropdownDefaultButton"
            data-dropdown-toggle="dropdown"
            className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 mr-1"
            type="button"
          >
            <RiEqualizerFill />
          </button>

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
      </div>
      {/* Card Course */}
      <div className="grid grid-cols-1 md:grid-cols-3 mt-5">
        {courses.map((item) => (
          <CourseCard
            key={item.id}
            id={item.id}
            name={item.name}
            description={item.teacher_id}
            image={item.image}
          />
        ))}
      </div>
    </div>
  );
};

export default CourseList;
