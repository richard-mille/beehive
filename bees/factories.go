/*
 *    Copyright (C) 2014-2017 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

// Package bees is Beehive's central module system.
package bees

import (
	"errors"

	"golang.org/x/oauth2"
)

// A BeeFactory is the base struct to be embedded by other BeeFactories.
type BeeFactory struct {
}

// Image returns an empty image filename per default.
func (factory *BeeFactory) Image() string {
	return ""
}

// LogoColor returns the default logo color.
func (factory *BeeFactory) LogoColor() string {
	return "#010000"
}

// OAuth2AccessToken returns the oauth2 access token.
func (factory *BeeFactory) OAuth2AccessToken(id, secret, code string) (*oauth2.Token, error) {
	return nil, errors.New("This Hive does not implement OAuth2")
}

// Options returns the default empty options set.
func (factory *BeeFactory) Options() []BeeOptionDescriptor {
	return []BeeOptionDescriptor{}
}

// States returns the default empty states set.
func (factory *BeeFactory) States() []StateDescriptor {
	return []StateDescriptor{}
}

// Events returns the default empty events set.
func (factory *BeeFactory) Events() []EventDescriptor {
	return []EventDescriptor{}
}

// Actions returns the default empty actions set.
func (factory *BeeFactory) Actions() []ActionDescriptor {
	return []ActionDescriptor{}
}

// A BeeFactoryInterface is the interface that gets implemented by a BeeFactory.
type BeeFactoryInterface interface {
	// ID of the module
	ID() string
	// Name of the module
	Name() string
	// Description of the module
	Description() string
	// An image url for the module
	Image() string
	// A logo color for the module
	LogoColor() string

	// OAuth2AccessToken returns the oauth2 access token.
	OAuth2AccessToken(id, secret, code string) (*oauth2.Token, error)

	// Options supported by module
	Options() []BeeOptionDescriptor
	// States provided by module
	States() []StateDescriptor
	// Events defined by module
	Events() []EventDescriptor
	// Actions supported by module
	Actions() []ActionDescriptor

	New(name, description string, options BeeOptions) BeeInterface
}

// RegisterFactory gets called by BeeFactories to register themselves.
func RegisterFactory(factory BeeFactoryInterface) {
	// log.Println("Bee Factory ready:", factory.ID(), "-", factory.Description())
	/* for _, ev := range factory.Events() {
		log.Println("\tProvides event:", ev.Name, "-", ev.Description)
		for _, opt := range ev.Options {
			log.Println("\t\tPlaceholder:", opt.Name, "-", opt.Description)
		}
	}
	for _, ac := range factory.Actions() {
		log.Println("\tOffers action:", ac.Name, "-", ac.Description)
		for _, opt := range ac.Options {
			log.Println("\t\tPlaceholder:", opt.Name, "-", opt.Description)
		}
	}
	log.Println() */

	factories[factory.ID()] = &factory
}

// GetFactory returns the factory with a specific name.
func GetFactory(identifier string) *BeeFactoryInterface {
	factory, ok := factories[identifier]
	if ok {
		return factory
	}

	return nil
}

// GetFactories returns all known bee factories.
func GetFactories() []*BeeFactoryInterface {
	r := []*BeeFactoryInterface{}
	for _, factory := range factories {
		r = append(r, factory)
	}

	return r
}
