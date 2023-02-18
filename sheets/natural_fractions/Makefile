# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make natural_fractions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

natural_fractions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_fractions.tex

natural_fractions.pdf: ../../*.tex ../../trademark.pdf *.tex natural_fractions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode natural_fractions.tex
	../../latexrun natural_fractions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_fractions.pdf
	open natural_fractions.pdf

o:
	make open

terms: natural_fractions.tex
	bbk_sheet -mode ts natural_fractions.tex

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
	rm -f natural_fractions.*
	make reset
	bbk_sheet -mode mk > Makefile