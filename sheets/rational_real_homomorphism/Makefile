# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make rational_real_homomorphism.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

rational_real_homomorphism.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > rational_real_homomorphism.tex

rational_real_homomorphism.pdf: ../../*.tex ../../trademark.pdf *.tex rational_real_homomorphism.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode rational_real_homomorphism.tex
	../../latexrun rational_real_homomorphism.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rational_real_homomorphism.pdf
	open rational_real_homomorphism.pdf

o:
	make open

terms: rational_real_homomorphism.tex
	bbk_sheet -mode ts rational_real_homomorphism.tex

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
	rm -f rational_real_homomorphism.*
	make reset
	bbk_sheet -mode mk > Makefile