"use client"

import { Button } from '@/components/ui/button';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { signInSchema } from '@/lib/zod';
import { useTheme } from 'next-themes';
import Image from 'next/image';
import React from 'react'
import { useForm } from 'react-hook-form';
import { FaMoon } from 'react-icons/fa';
import { MdOutlineWbSunny } from 'react-icons/md';
import { z } from 'zod';
import { zodResolver } from "@hookform/resolvers/zod";
import { signIn } from '../_actions/signin';
import { useRouter, useSearchParams } from 'next/navigation';
import LogoIpsum from '@/public/images/logoipsum-288.svg';

const LoginForm = () => {
    const { setTheme, theme } = useTheme();
    const router = useRouter()
    const searchParams = useSearchParams()

    const form = useForm<z.infer<typeof signInSchema>>({
        resolver: zodResolver(signInSchema),
        defaultValues: {
            email: "",
            password: "",
        },
    });

    const { isSubmitting, isValid } = form.formState;

    const onSubmit = async (values: z.infer<typeof signInSchema>) => {
        const { isLoggedIn, errorMessage } = await signIn(values)

        if (!isLoggedIn) {
            form.setError("password", {
                type: "validate",
                message: errorMessage || "Invalid credentials",
            })

            return
        }
        
        router.push(searchParams.get("callback") || "/")
    }

    return (
        <div className="bg-gradient-to-r dark:from-gray-900 dark:to-black">
            <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
                <a href="#" className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
                    <Image
                        src={LogoIpsum}
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
                        <Form {...form}>
                            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4 md:space-y-6">
                                <FormField
                                    control={form.control}
                                    name="email"
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormLabel className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                                                Your email
                                            </FormLabel>
                                            <FormControl>
                                                <Input
                                                    disabled={isSubmitting}
                                                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-600 focus:border-blue-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-x-blue-500"
                                                    {...field}
                                                />
                                            </FormControl>
                                            <FormMessage  />
                                        </FormItem>
                                    )} />
                                <FormField
                                    control={form.control}
                                    name="password"
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormLabel className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                                                Password
                                            </FormLabel>
                                            <FormControl>
                                                <Input
                                                    disabled={isSubmitting}
                                                    type="password"
                                                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-600 focus:border-blue-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                    {...field}
                                                />
                                            </FormControl>
                                            <FormMessage />
                                        </FormItem>
                                    )} />
                                {/* <FormField
                                    control={form.control}
                                    name="remember"
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormControl>
                                                <Checkbox
                                                    className="w-4 h-4 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-blue-600 dark:ring-offset-gray-800"
                                                    {...field}
                                                />
                                            </FormControl>
                                            <FormLabel
                                                className="ml-3 text-sm text-gray-500 dark:text-gray-300"
                                            >
                                                Remember me
                                            </FormLabel>
                                        </FormItem>
                                    )} /> */}
                                {/* <a
                                    href="#"
                                    className="text-sm font-medium text-white hover:underline"
                                >
                                    Forgot password?
                                </a> */}
                                <Button
                                    disabled={isSubmitting}
                                    type="submit"
                                    className="w-full text-white font-medium rounded-lg text-sm px-5 py-2.5 text-center transition ease-in delay-150 hover:-translate-y-1 hover:scale-110 hover:bg-gray-900 duration-200 bg-gray-700 active:bg-gray-900 focus:outline-none focus: ring focus:ring-gray-300">
                                    Login
                                </Button>
                            </form>
                        </Form>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default LoginForm
