"use client"

import { useTheme } from 'next-themes';
import Image from 'next/image';
import React, { useState } from 'react'
import { FaMoon } from 'react-icons/fa';
import { MdOutlineWbSunny } from 'react-icons/md';

const LoginForm = () => {
    const { setTheme, theme } = useTheme();
    const [isFailed, setIsFailed] = useState(true);

    return (
        <div className="bg-gradient-to-r from-gray-900 to-black">
            <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
                <a href="#" className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
                    <Image
                        src="/images/logoipsum-288.svg"
                        alt=""
                        width={"100"}
                        height={"40"}
                    />
                </a>
                <div className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
                    <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
                        <div className="flex justify-between">
                            <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
                                Sign in to your account
                            </h1>
                            <button
                                className="w-8 h-8"
                                onClick={() => {
                                    setTheme(theme === "dark" ? "light" : "dark");
                                }}>

                                <FaMoon className="flex-1 text-white w-8 h-8 hover:bg-white p-1 hover:rounded-lg hover:text-black hidden dark:block" />
                                <MdOutlineWbSunny className="flex-1 text-black w-8 h-8 p-0.5 hover:bg-black hover:rounded-lg hover:text-white dark:hidden" />
                            </button>
                        </div>
                        <form action="#" className="space-y-4 md:space-y-6">
                            <div>
                                <label
                                    htmlFor="email"
                                    className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                                >
                                    Your email
                                </label>
                                <input
                                    type="email"
                                    name="email"
                                    id="email"
                                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-600 focus:border-blue-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-x-blue-500"
                                    placeholder="email@gmail.com"
                                    required
                                />
                                <span
                                    className={`text-xs italic text-red-700 pl-1 ${isFailed ? "" : "hidden"
                                        }`}
                                >
                                    Email is not correct!
                                </span>
                            </div>
                            <div>
                                <label
                                    htmlFor="password"
                                    className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                                >
                                    Password
                                </label>
                                <input
                                    type="password"
                                    name="password"
                                    id="password"
                                    placeholder="Your Password"
                                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-600 focus:border-blue-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                    required
                                />
                                <span className={`text-xs italic text-red-700 pl-1 ${isFailed ? "" : "hidden"}`}>
                                    Password is not correct!
                                </span>
                            </div>
                            <div className="flex items-center justify-between">
                                <div className="flex items-start">
                                    <div className="flex items-center h-5">
                                        <input
                                            type="checkbox"
                                            id="remember"
                                            className="w-4 h-4 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-blue-600 dark:ring-offset-gray-800"
                                            required
                                        />
                                    </div>
                                    <div className="ml-3 text-sm">
                                        <label
                                            htmlFor="remember"
                                            className="text-gray-500 dark:text-gray-300"
                                        >
                                            Remember me
                                        </label>
                                    </div>
                                </div>
                                <a
                                    href="#"
                                    className="text-sm font-medium text-white hover:underline"
                                >
                                    Forgot password?
                                </a>
                            </div>
                            <button
                                onClick={() => {
                                    setIsFailed((fail) => !fail);
                                }}
                                className="w-full text-white font-medium rounded-lg text-sm px-5 py-2.5 text-center transition ease-in delay-150 hover:-translate-y-1 hover:scale-110 hover:bg-gray-900 duration-200 bg-gray-700 active:bg-gray-900 focus:outline-none focus: ring focus:ring-gray-300"
                            >
                                Login
                            </button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default LoginForm
