<?php

namespace App\Http\Livewire\Front;

use App\Models\ContactSubmission;
use Artesaos\SEOTools\Facades\SEOTools;
use DanHarrin\LivewireRateLimiting\Exceptions\TooManyRequestsException;
use DanHarrin\LivewireRateLimiting\WithRateLimiting;
use Livewire\Component;

class Contact extends Component
{
    use WithRateLimiting;

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

    public function mount()
    {
        SEOTools::setTitle('Contact');
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

        try {
            $this->rateLimit(3);
        } catch (TooManyRequestsException $exception) {
            $this->addError('data.email', "Slow down! Please wait another $exception->secondsUntilAvailable seconds to log in.");
            return;
        }

        ContactSubmission::create($this->data);

        $this->reset();

        session()->flash('message', 'Your message has been sent!');
    }
}
