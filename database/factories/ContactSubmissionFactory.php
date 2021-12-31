<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

class ContactSubmissionFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array
     */
    public function definition()
    {
        return [
            'email' => $this->faker->email,
            'message' => $this->faker->sentence,
            'subject' => $this->faker->sentence,
            'name' => $this->faker->name,
        ];
    }
}
