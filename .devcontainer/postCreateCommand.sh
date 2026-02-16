#!/bin/bash

# postCreateCommand.sh
# This script runs after the dev container is created
# It installs dependencies for Ruby, Node.js, and Go if the respective files exist

set -e

echo "🚀 Running post-create setup..."

# Install Ruby dependencies if Gemfile exists
if [ -f "Gemfile" ]; then
    echo "📦 Installing Ruby dependencies (bundle install)..."
    bundle install
else
    echo "⏭️  No Gemfile found, skipping Ruby dependencies"
fi

# Install Node.js dependencies if package.json exists
if [ -f "package.json" ]; then
    echo "📦 Installing Node.js dependencies (npm install)..."
    npm install
else
    echo "⏭️  No package.json found, skipping Node.js dependencies"
fi

# Install Go dependencies if go.mod exists
if [ -f "go.mod" ]; then
    echo "📦 Installing Go dependencies (go mod download)..."
    go mod download
else
    echo "⏭️  No go.mod found, skipping Go dependencies"
fi

echo "✅ Post-create setup complete!"
