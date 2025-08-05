# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make group_generators.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

group_generators.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > group_generators.tex

group_generators.pdf: ../../*.tex ../../trademark.pdf *.tex group_generators.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode group_generators.tex
	../../latexrun group_generators.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: group_generators.pdf
	open group_generators.pdf

o:
	make open

terms: group_generators.tex
	bbk sheet -mode ts group_generators.tex

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
	rm -f group_generators.*
	make reset
	bbk sheet -mode mk > Makefile