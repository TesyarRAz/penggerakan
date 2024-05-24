import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import React from 'react'
import {Control, useFormContext, UseFormReturn} from 'react-hook-form'

const ImageForm = () => {
    const form = useFormContext()

    const { isSubmitting } = form.formState

    return (
        <FormField
            control={form.control}
            name="image"
            render={({ field }) => (
                <FormItem>
                    <FormLabel>
                        Image URL
                    </FormLabel>
                    <FormControl>
                        <Input
                            inputMode="url"
                            disabled={isSubmitting}
                            {...field}
                        />
                    </FormControl>
                    <FormMessage />
                </FormItem>
            )}
        />
    )
}

export default ImageForm
