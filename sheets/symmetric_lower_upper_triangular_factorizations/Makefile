# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make symmetric_lower_upper_triangular_factorizations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

symmetric_lower_upper_triangular_factorizations.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > symmetric_lower_upper_triangular_factorizations.tex

symmetric_lower_upper_triangular_factorizations.pdf: ../../*.tex ../../trademark.pdf *.tex symmetric_lower_upper_triangular_factorizations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode symmetric_lower_upper_triangular_factorizations.tex
	../../latexrun symmetric_lower_upper_triangular_factorizations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: symmetric_lower_upper_triangular_factorizations.pdf
	open symmetric_lower_upper_triangular_factorizations.pdf

o:
	make open

terms: symmetric_lower_upper_triangular_factorizations.tex
	bbk_sheet -mode ts symmetric_lower_upper_triangular_factorizations.tex

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
	rm -f symmetric_lower_upper_triangular_factorizations.*
	make reset
	bbk_sheet -mode mk > Makefile