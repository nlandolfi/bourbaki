# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make topological_groups.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

topological_groups.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > topological_groups.tex

topological_groups.pdf: ../../*.tex ../../trademark.pdf *.tex topological_groups.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode topological_groups.tex
	../../latexrun topological_groups.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: topological_groups.pdf
	open topological_groups.pdf

o:
	make open

terms: topological_groups.tex
	bbk sheet -mode ts topological_groups.tex

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
	rm -f topological_groups.*
	make reset
	bbk sheet -mode mk > Makefile