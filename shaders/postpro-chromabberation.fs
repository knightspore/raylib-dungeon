#version 410 core

in vec2 fragTexCoord;
in vec4 fragColor;
out vec4 finalColor;
uniform sampler2D tex;

uniform float chromaticAberration = 0.005;
uniform vec2 center = vec2(0.5, 0.5);

void main() {
    float dist = length(fragTexCoord - center);

    // Clamp to prevent screen wrapping?

    vec4 r = texture(tex, fragTexCoord + vec2(chromaticAberration * dist, 0.0));
    vec4 g = texture(tex, fragTexCoord);
    vec4 b = texture(tex, fragTexCoord - vec2(chromaticAberration * dist, 0.0));

    finalColor = vec4(r.r, g.g, b.b, 1.0);
}
