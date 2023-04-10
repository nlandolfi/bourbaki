# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make algebraic_equations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

algebraic_equations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > algebraic_equations.tex

algebraic_equations.pdf: ../../*.tex ../../trademark.pdf *.tex algebraic_equations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode algebraic_equations.tex
	../../latexrun algebraic_equations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: algebraic_equations.pdf
	open algebraic_equations.pdf

o:
	make open

terms: algebraic_equations.tex
	bbk_sheet -mode ts algebraic_equations.tex

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
	rm -f algebraic_equations.*
	make reset
	bbk_sheet -mode mk > Makefile