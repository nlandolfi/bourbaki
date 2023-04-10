# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make topology_bases.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

topology_bases.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > topology_bases.tex

topology_bases.pdf: ../../*.tex ../../trademark.pdf *.tex topology_bases.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode topology_bases.tex
	../../latexrun topology_bases.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: topology_bases.pdf
	open topology_bases.pdf

o:
	make open

terms: topology_bases.tex
	bbk sheet -mode ts topology_bases.tex

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
	rm -f topology_bases.*
	make reset
	bbk sheet -mode mk > Makefile