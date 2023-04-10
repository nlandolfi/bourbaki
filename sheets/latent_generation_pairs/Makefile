# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make latent_generation_pairs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

latent_generation_pairs.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > latent_generation_pairs.tex

latent_generation_pairs.pdf: ../../*.tex ../../trademark.pdf *.tex latent_generation_pairs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode latent_generation_pairs.tex
	../../latexrun latent_generation_pairs.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: latent_generation_pairs.pdf
	open latent_generation_pairs.pdf

o:
	make open

terms: latent_generation_pairs.tex
	bbk sheet -mode ts latent_generation_pairs.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f latent_generation_pairs.*
	make reset
	bbk sheet -mode mk > Makefile