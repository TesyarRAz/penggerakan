import React from "react";
import CrudProfileCard from "../components/profile-card";
import ProfileCard from "../components/profile-card";
import { FaSearch } from "react-icons/fa";
import { RiEqualizerFill } from "react-icons/ri";
import { IoAddCircleOutline } from "react-icons/io5";

const studentProfile = [
  {
    name: "Teacher1",
    grade: "XI-A",
    major: "Computer Science",
    image: "/images/cat-4.jpg",
  },
  {
    name: "Teacher2",
    grade: "XI-A",
    major: "Computer Science",
    image: "/images/cat-4.jpg",
  },
  {
    name: "Teacher3",
    grade: "XI-A",
    major: "Computer Science",
    image: "/images/cat-4.jpg",
  },
  {
    name: "Teacher4",
    grade: "XI-A",
    major: "Computer Science",
    image: "/images/cat-4.jpg",
  },
  {
    name: "Teacher5",
    grade: "XI-A",
    major: "Computer Science",
    image: "/images/cat-4.jpg",
  },
  {
    name: "Teacher6",
    grade: "XI-A",
    major: "Computer Science",
    image: "/images/cat-4.jpg",
  },
];

const TeacherAdmin = () => {
  return (
    <div className="bg-white dark:bg-gray-900 h-screen overflow-auto scrollbar-thin dark:scrollbar-thin scroll-smooth">
      <div className="bg-gray-200 mt-5 mx-5 p-2 rounded-lg dark:bg-gray-800">
        <div className="flex justify-between">
          <h2 className="font-sans font-bold text-3xl ml-5 dark:text-white">
            Student List
          </h2>
          <div className="mr-3">
            <button
              type="button"
              className="text-green-700 hover:text-white border border-green-700 hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-green-500 dark:text-green-500 dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-green-800"
            >
              <div className="flex justify-center items-center">
                <IoAddCircleOutline className="mr-1 h-5 w-5" />
                <span>Add Student</span>
              </div>
            </button>
          </div>
        </div>
        <form action="#" className="flex mt-5">
          <input
            type="text"
            id="browseStudent"
            name="browseStudent"
            placeholder="Browse Student..."
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
              type="button"
              className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm  text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 p-3 mr-1"
            >
              <RiEqualizerFill />
            </button>
          </div>
        </form>
        <div className="grid grid-cols-1 md:grid-cols-3 ml-2">
          <ProfileCard students={studentProfile} />
        </div>
      </div>
    </div>
  );
};

export default TeacherAdmin;
