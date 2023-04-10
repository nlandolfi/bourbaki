# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make natural_additive_identity.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

natural_additive_identity.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > natural_additive_identity.tex

natural_additive_identity.pdf: ../../*.tex ../../trademark.pdf *.tex natural_additive_identity.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode natural_additive_identity.tex
	../../latexrun natural_additive_identity.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_additive_identity.pdf
	open natural_additive_identity.pdf

o:
	make open

terms: natural_additive_identity.tex
	bbk sheet -mode ts natural_additive_identity.tex

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
	rm -f natural_additive_identity.*
	make reset
	bbk sheet -mode mk > Makefile