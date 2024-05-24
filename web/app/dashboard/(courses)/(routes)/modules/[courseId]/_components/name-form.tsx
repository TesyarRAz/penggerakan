import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import React from 'react'
import {Control, useFormContext, UseFormReturn} from 'react-hook-form'

const NameForm = () => {
    const form = useFormContext()

    const { isSubmitting } = form.formState

    return (
        <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
                <FormItem>
                    <FormLabel>
                        Name
                    </FormLabel>
                    <FormControl>
                        <Input
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

export default NameForm
