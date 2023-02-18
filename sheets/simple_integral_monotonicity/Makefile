# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make simple_integral_monotonicity.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

simple_integral_monotonicity.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > simple_integral_monotonicity.tex

simple_integral_monotonicity.pdf: ../../*.tex ../../trademark.pdf *.tex simple_integral_monotonicity.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode simple_integral_monotonicity.tex
	../../latexrun simple_integral_monotonicity.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: simple_integral_monotonicity.pdf
	open simple_integral_monotonicity.pdf

o:
	make open

terms: simple_integral_monotonicity.tex
	bbk_sheet -mode ts simple_integral_monotonicity.tex

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
	rm -f simple_integral_monotonicity.*
	make reset
	bbk_sheet -mode mk > Makefile