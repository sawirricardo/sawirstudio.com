<?php

namespace App\Http\Livewire\Front;

use App\Models\ContactSubmission;
use Livewire\Component;

class Contact extends Component
{
    public $data = [
        'name' => '',
        'email' => '',
        'subject' => '',
        'message' => '',
    ];

    protected function rules()
    {
        return [
            'data.name' => ['string'],
            'data.subject' => ['string'],
            'data.email' => ['required', 'email'],
            'data.message' => ['required'],
        ];
    }

    public function render()
    {
        return view('livewire.front.contact');
    }

    public function updated($propertyName)
    {
        $this->validateOnly($propertyName);
    }

    public function create()
    {
        $this->validate();

        ContactSubmission::create($this->data);

        $this->reset();

        session()->flash('message', 'Your message has been sent!');
    }
}
