# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make natural_powers.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

natural_powers.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > natural_powers.tex

natural_powers.pdf: ../../*.tex ../../trademark.pdf *.tex natural_powers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode natural_powers.tex
	../../latexrun natural_powers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_powers.pdf
	open natural_powers.pdf

o:
	make open

terms: natural_powers.tex
	bbk_sheet -mode ts natural_powers.tex

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
	rm -f natural_powers.*
	make reset
	bbk_sheet -mode mk > Makefile