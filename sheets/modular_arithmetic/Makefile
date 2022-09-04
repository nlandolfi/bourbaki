# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make modular_arithmetic.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

modular_arithmetic.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > modular_arithmetic.tex

modular_arithmetic.pdf: ../../*.tex ../../trademark.pdf *.tex modular_arithmetic.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode modular_arithmetic.tex
	../../latexrun modular_arithmetic.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: modular_arithmetic.pdf
	open modular_arithmetic.pdf

o:
	make open

terms: modular_arithmetic.tex
	bbk_sheet -mode ts modular_arithmetic.tex

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
	rm -f modular_arithmetic.*
	make reset
	bbk_sheet -mode mk > Makefile