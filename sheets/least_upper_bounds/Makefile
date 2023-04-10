# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make least_upper_bounds.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

least_upper_bounds.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > least_upper_bounds.tex

least_upper_bounds.pdf: ../../*.tex ../../trademark.pdf *.tex least_upper_bounds.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode least_upper_bounds.tex
	../../latexrun least_upper_bounds.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: least_upper_bounds.pdf
	open least_upper_bounds.pdf

o:
	make open

terms: least_upper_bounds.tex
	bbk_sheet -mode ts least_upper_bounds.tex

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
	rm -f least_upper_bounds.*
	make reset
	bbk_sheet -mode mk > Makefile